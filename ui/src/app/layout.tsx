import type { Metadata } from "next";
import {
  ColorSchemeScript,
  createTheme,
  mantineHtmlProps,
  MantineProvider,
} from "@mantine/core";
import { ReactNode } from "react";

import "@mantine/core/styles.css";
import "@mantine/notifications/styles.css";
import "./globals.css";
import { Instrument_Sans, Instrument_Serif } from "next/font/google";

const bodyFont = Instrument_Sans({
  weight: ["400", "500", "600", "700"],
  subsets: ["latin"],
  display: "swap",
  variable: "--body-font",
});

const titleFont = Instrument_Serif({
  weight: ["400"],
  subsets: ["latin"],
  display: "swap",
  variable: "--title-font",
});

export const metadata: Metadata = {
  title: "Code of Federal Regulations - Metrics",
  description: "Enhanced analysis of Federal Regulations data",
};

const theme = createTheme({
  colors: {
    primary: [
      "var(--primary-shade-50)",
      "var(--primary-shade-100)",
      "var(--primary-shade-200)",
      "var(--primary-shade-300)",
      "var(--primary-shade-400)",
      "var(--primary-shade-500)",
      "var(--primary-shade-600)",
      "var(--primary-shade-700)",
      "var(--primary-shade-800)",
      "var(--primary-shade-900)",
    ],
    emphasis: [
      "var(--emphasis-shade-50)",
      "var(--emphasis-shade-100)",
      "var(--emphasis-shade-200)",
      "var(--emphasis-shade-300)",
      "var(--emphasis-shade-400)",
      "var(--emphasis-shade-500)",
      "var(--emphasis-shade-600)",
      "var(--emphasis-shade-700)",
      "var(--emphasis-shade-800)",
      "var(--emphasis-shade-900)",
    ],
  },
  primaryColor: "primary",
});

export default function RootLayout({
  children,
}: Readonly<{
  children: ReactNode;
}>) {
  return (
    <html
      lang="en"
      {...mantineHtmlProps}
      className={`${bodyFont.variable} ${titleFont.variable}`}
    >
      <head>
        <link
          rel="icon"
          href="data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22><text y=%22.9em%22 font-size=%2290%22>🇺🇸</text></svg>"
        />
        <ColorSchemeScript />
      </head>
      <body>
        <MantineProvider theme={theme}>{children}</MantineProvider>
      </body>
    </html>
  );
}
