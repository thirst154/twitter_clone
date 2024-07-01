import { Heart, Pencil, Trash2 } from "lucide-react";
import { Card, CardContent, CardFooter, CardHeader } from "./ui/card";
import { ConfirmDeleteDialog } from "./ConfirmDeleteDialog";
import { Button } from "./ui/button";
import { Post } from "@/lib/types";

export default function Message(post: Post) {
  const { id, text, date_published, user } = post || {};
  return (
    <div className="flex flex-col rounded-lg border bg-background p-3">
      <div className="flex justify-between">
        <div className="flex items-center gap-2">
          <div className="w-10 h-10 bg-primary rounded-full border"></div>
          <div className="grid max-w-fit items-center">
            <span className="text-lg font-semibold">
              {user.first_name} {user.last_name}
            </span>
            <span className="text-sm text-muted-foreground">
              {user.username}
            </span>
          </div>
        </div>
        <Button variant="outline">Follow</Button>
      </div>
      <div>
        <span className="text-sm">{text}</span>
      </div>
      <div>
        <div className="flex w-full justify-around gap-2 ">
          <Button variant="ghost" className="flex gap-2">
            <Heart size={20} /> Like
          </Button>
          <Button variant="ghost" className="flex gap-2">
            <Pencil size={20} />
            Edit
          </Button>
          <ConfirmDeleteDialog />
        </div>
      </div>
    </div>
  );
}
