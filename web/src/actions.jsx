import { redirect } from "react-router-dom";

// const backend = import.meta.env.VITE_BACKEND_URL;
const backend = "/api"

export async function action({ request, params }) {
  const token = localStorage.getItem("token");

  switch (request.method) {
    case "POST": {
      let formData = await request.formData();
      let title = formData.get("title");
      let body = formData.get("body");
      await fetch(`${backend}/blog`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ title, body }),
      });

      return redirect("/dashboard");
    }
    case "PUT": {
      let formData = await request.formData();
      let title = formData.get("title");
      let body = formData.get("body");
      await fetch(`${backend}/blog/${params.id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "appication/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ title, body }),
      });

      return redirect(`/dashboard`);
    }

    case "DELETE": {
      await fetch(`${backend}/blog/${params.id}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      return redirect("/dashboard", { replace: true });
    }
  }
}

export async function loginAction({ request }) {
  let formData = await request.formData();
  let username = formData.get("username");
  let password = formData.get("password");
  const res = await fetch(`${backend}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username, password }),
  });
  if (res.ok) {
    const data = await res.json();
    localStorage.setItem("token", data.token);
    return redirect("/dashboard", { replace: true });
  } else {
    alert("Wrong Username or Password");
    return redirect("/login");
  }
}
