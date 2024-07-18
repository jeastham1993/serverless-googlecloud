"use client";

import "@fontsource/roboto/300.css";
import "@fontsource/roboto/400.css";
import "@fontsource/roboto/500.css";
import "@fontsource/roboto/700.css";
import "../public/style.css";

import { Container } from "@mui/material";
import MenuAppBar from "./components/menu-app-bar";
import { initializeApp } from "firebase/app";
import { getAuth, createUserWithEmailAndPassword } from "firebase/auth";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <MenuAppBar />
        <Container maxWidth="xl">{children}</Container>
      </body>
    </html>
  );
}
