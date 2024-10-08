import { Link, useLoaderData } from "react-router-dom";
import "./BlogPageList.css";
import Time from "../time";

export default function BlogPageList() {
  const blogs = useLoaderData();

  return (
    <section>
      <header>
        <h1>List of Blogs</h1>
      </header>
      {blogs ? <BlogList blogs={blogs} /> : <p>No blogs</p>}
    </section>
  );
}

function BlogList({ blogs }) {
  return (
    <ul className='bloglist'>
      {blogs.map((blog) => (
        <li key={blog.id}>
          <Link to={`/blogs/${blog.id}`}>{blog.title}</Link>
          <Time blog={blog} />
        </li>
      ))}
    </ul>
  );
}
