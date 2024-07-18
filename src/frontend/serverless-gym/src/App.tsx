import React from "react";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import { Box, CssBaseline } from "@mui/material";
import "./App.css";
import Home from "./components/home";
import Session from "./components/session";
import ViewSession from "./components/viewSession";
import MenuAppBar from "./components/menu-app-bar";
import Login from "./components/login";
import { datadogRum } from "@datadog/browser-rum";

datadogRum.init({
  applicationId: "68878d41-e394-4eca-93ee-f3a979c412bb",
  clientToken: "pub10900ce49e775e074776d40f1b08576c",
  // `site` refers to the Datadog site parameter of your organization
  // see https://docs.datadoghq.com/getting_started/site/
  site: "datadoghq.eu",
  service: "serverless-gym",
  env: "dev",
  // Specify a version number to identify the deployed version of your application in Datadog
  // version: '1.0.0',
  sessionSampleRate: 100,
  sessionReplaySampleRate: 20,
  trackUserInteractions: true,
  trackResources: true,
  trackLongTasks: true,
  defaultPrivacyLevel: "mask-user-input",
});

function App() {
  return (
    <div className="App">
      <Router>
        <MenuAppBar />
        <Box sx={{ flexGrow: 1 }}>
          <CssBaseline />
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
            <Route path="/session" element={<Session />} />
            <Route path="/session/:id" element={<ViewSession />} />
          </Routes>
        </Box>
      </Router>
    </div>
  );
}

export default App;
