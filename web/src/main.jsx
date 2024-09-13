import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./App.jsx";
import "./index.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import HomePage from "./pages/Homepage/HomePage.jsx";
import ErrorPage from "./pages/error-page.jsx";
import BlogForm from "./pages/Create Blog Page/BlogForm.jsx";
import Blogs from "./pages/Blogs Page/Blogs.jsx";
import Blog from "./pages/Blog Page/Blog.jsx";
import UpdateBlogPage from "./pages/Update Blog Page/UpdateBlogPage.jsx";
import Dashboard from "./pages/Dashboard/Dashboard.jsx";
import { action, loginAction } from "./actions.jsx";
import { getBlogLoader } from "./loaders/getBlogLoader.jsx";
import { getBlogsLoader } from "./loaders/getBlogsLoader.jsx";
import Login from "./pages/Login Page/Login.jsx";
import ProtectedRoute from "./components/ProtectedRoute/ProtectedRoute.jsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "/",
        element: <HomePage />,
        loader: getBlogsLoader,
      },
      {
        path: "blogs",
        element: <Blogs />,
        loader: getBlogsLoader,
      },
      {
        path: "blogs/:id",
        element: <Blog />,
        action: action,
        loader: getBlogLoader,
      },
      {
        path: "dashboard",
        element: (
          <ProtectedRoute>
            <Dashboard />
          </ProtectedRoute>
        ),
        loader: getBlogsLoader,
        action: action,
      },
      {
        path: "dashboard/create",
        element: (
          <ProtectedRoute>
            <BlogForm />
          </ProtectedRoute>
        ),
        action: action,
      },
      {
        path: "dashboard/:id",
        action: action,
      },
      {
        path: "dashboard/:id/update",
        element: (
          <ProtectedRoute>
            <UpdateBlogPage />
          </ProtectedRoute>
        ),
        loader: getBlogLoader,
        action: action,
      },
      {
        path: "login",
        element: <Login />,
        action: loginAction,
      },
    ],
  },
]);

createRoot(document.getElementById("root")).render(
  // <StrictMode>
  <RouterProvider router={router} />
  // </StrictMode>
);
