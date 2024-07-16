using GoogleCloudRunSample.Product;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Npgsql;

var builder = WebApplication.CreateBuilder(args);
builder.Configuration.AddEnvironmentVariables();

builder.Services.AddDbContext<ProductDbContext>(options =>
    options.UseNpgsql(GetPostgreSqlConnectionString(builder.Configuration).ToString()));

var app = builder.Build();

using (var scope = app.Services.CreateScope())
{
    var context = scope.ServiceProvider.GetRequiredService<ProductDbContext>();
    context.Database.EnsureCreated();
}

app.MapGet("/product/{productId}", (ProductDbContext context, string productId) =>
{
    var product = context.Products.FirstOrDefault(p => p.ProductId == productId);

    if (product is null)
    {
        return Results.NotFound();
    }

    return Results.Ok(product);
});

app.MapPost("/product", async (ProductDbContext context, [FromBody] Product product) =>
{
    product.ProductId = Guid.NewGuid().ToString();
    
    context.Products.Add(product);
    await context.SaveChangesAsync();

    return product;
});

app.Run();

static NpgsqlConnectionStringBuilder GetPostgreSqlConnectionString(IConfiguration configuration)
{
    // Allow a full connection string to be passed in, useful for local development
    if (configuration["DB_CONNECTION_STRING"] != null)
    {
        return new NpgsqlConnectionStringBuilder(configuration["DB_CONNECTION_STRING"]);
    }
    
    // Failing that, use a UNIX socket connection for running inside GoogleCloudRun and using Cloud SQL
    var connectionString = new NpgsqlConnectionStringBuilder()
    {
        SslMode = SslMode.Disable,
        Host = configuration["INSTANCE_UNIX_SOCKET"], // e.g. '/cloudsql/project:region:instance'
        Username = configuration["DB_USER"], // e.g. 'my-db-user
        Password = configuration["DB_PASS"], // e.g. 'my-db-password'
        Database = configuration["DB_NAME"], // e.g. 'my-database'
    };
    
    connectionString.Pooling = true;
    connectionString.MaxPoolSize = 5;
    connectionString.MinPoolSize = 0;
    connectionString.Timeout = 15;
    connectionString.ConnectionIdleLifetime = 300;
    return connectionString;
}