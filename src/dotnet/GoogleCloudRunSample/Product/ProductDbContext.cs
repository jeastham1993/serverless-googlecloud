using Microsoft.EntityFrameworkCore;

namespace GoogleCloudRunSample.Product;

public class ProductDbContext : DbContext
{
    public ProductDbContext(DbContextOptions options) : base(options)
    {
    }
    
    public DbSet<Product> Products { get; set; }
}