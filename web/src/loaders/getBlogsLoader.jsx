export async function getBlogsLoader() {
  // const backend = import.meta.env.VITE_BACKEND_URL;
  const backend = "/api"
  const res = await fetch(`${backend}/blog`);
  if (res.ok) {
    const blogs = await res.json();
    return blogs;
  }
}
