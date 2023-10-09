import * as React from "react";
import * as ReactDOM from "react-dom/client";
import CssBaseline from "@mui/material/CssBaseline";
import { ThemeProvider } from "@mui/material/styles";
import theme from "./theme";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Home from "./pages/Home";
import CategoryPage from "./pages/CategoryPage";
import CreateCategoryPage from "./pages/CreateCategoryPage";
import LoginPage from "./pages/LoginPage";
import UpdateCategoryPage from "./pages/UpdateCategoryPage";
import ComplaintPage from "./pages/ComplaintPage";
import ComplaintDetailPage from "./pages/ComplaintDetailPage";
import CreateComplaintPage from "./pages/CreateComplaintPage";

const rootElement = document.getElementById("root");
const root = ReactDOM.createRoot(rootElement);

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/login",
    element: <LoginPage />,
  },
  {
    path: "/category",
    element: <CategoryPage />,
  },
  {
    path: "/category/create",
    element: <CreateCategoryPage />,
  },
  {
    path: "/category/:id/update",
    element: <UpdateCategoryPage />,
  },
  {
    path: "/complaint",
    element: <ComplaintPage />,
  },
  {
    path: "/complaint/create",
    element: <CreateComplaintPage />,
  },
  {
    path: "/complaint/:id",
    element: <ComplaintDetailPage />,
  },
]);

root.render(
  <ThemeProvider theme={theme}>
    <CssBaseline />
    <RouterProvider router={router} />
  </ThemeProvider>
);
