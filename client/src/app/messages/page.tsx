import Message from "@/components/Message";
import NewMessageForm from "@/components/NewMessageForm";
import { GetPostsData, Post } from "@/lib/types";

async function getPosts() {
  const res = await fetch("http://localhost:8080/posts", { cache: "no-store" });
  const data = await res.json();
  return data?.data as Post[];
}

export default async function Home() {
  const posts = await getPosts();
  return (
    <div className="mx-auto w-full max-w-screen-lg grid gap-4">
      <h1>Messages</h1>
      <NewMessageForm></NewMessageForm>
      {posts?.map((post: Post) => {
        return (
          <Message
            key={post.id}
            id={post.id}
            text={post.text}
            date_published={post.date_published}
            user={post.user}
          />
        );
      })}
    </div>
  );
}
