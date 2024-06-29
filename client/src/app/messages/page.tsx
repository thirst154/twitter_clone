import Message from "@/components/Message";
import NewMessageForm from "@/components/NewMessageForm";

export default function Home(){
  return (
    <div className="mx-auto w-full max-w-screen-lg grid gap-4">
      <h1>Hello Messages</h1>
      <NewMessageForm></NewMessageForm>
      <Message></Message>
    </div>
  )
}