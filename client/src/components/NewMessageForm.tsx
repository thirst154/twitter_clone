import Link from "next/link";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "./ui/card";
import { Button } from "./ui/button";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Textarea } from "./ui/textarea";

export default function NewMessageForm() {
  return (
    <Card className="mx-auto w-full">
      <CardHeader>
        <CardTitle className="text-xl">Send New Message</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="grid w-full gap-1.5">
          <Label htmlFor="message">Your message</Label>
          <Textarea placeholder="Type your message here." id="message" />
          <Button type="submit">Send</Button>
        </div>
      </CardContent>
    </Card>
  );
}
