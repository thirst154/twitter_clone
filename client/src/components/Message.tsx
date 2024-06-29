import { Heart, Pencil, Trash2 } from "lucide-react";
import { Card, CardContent, CardFooter, CardHeader } from "./ui/card";
import { ConfirmDeleteDialog } from "./ConfirmDeleteDialog";
import { Button } from "./ui/button";

export default function Message() {
  return (
    <div className="flex flex-col rounded-lg border bg-background p-3">
      <div className="flex justify-between">
        <div className="flex items-center gap-2">
          <div className="w-10 h-10 bg-red-400 rounded-full border"></div>
          <div className="grid max-w-fit items-center">
            <span className="text-lg font-semibold">Thomas Hirst</span>
            <span className="text-sm text-muted-foreground">@username</span>
          </div>
        </div>
        <Button variant="outline">Follow</Button>
      </div>
      <div>
        <span className="text-sm">
          Lorem ipsum dolor sit, amet consectetur adipisicing elit. A minus
          consequatur culpa officiis provident dolor recusandae, explicabo
          quibusdam optio animi corrupti alias cupiditate iure ipsa ratione
          voluptate itaque non quisquam!
        </span>
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
