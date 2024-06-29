import { Pencil, Trash2 } from "lucide-react";
import { Card, CardContent, CardFooter, CardHeader } from "./ui/card";
import { ConfirmDeleteDialog } from "./ConfirmDeleteDialog";
import { Button } from "./ui/button";

export default function Message(){

  return (
    <div className="flex flex-col rounded-lg border bg-background p-3">
      <div>
      <div className="grid max-w-fit items-center">
        <span className="text-lg font-semibold">Thomas Hirst</span>
        <span className="text-sm text-muted-foreground">@username</span>
      </div>
      </div>
      <div>
      <span className="text-sm">Lorem ipsum dolor sit, amet consectetur adipisicing elit. A minus consequatur culpa officiis provident dolor recusandae, explicabo quibusdam optio animi corrupti alias cupiditate iure ipsa ratione voluptate itaque non quisquam!</span>
      </div>
      <div>
        <div className="flex w-full justify-end gap-2 ">
        <Button variant="ghost"><Pencil size={20} /></Button>
        <ConfirmDeleteDialog/> 
        </div>
      </div>
    </div>
  )
}

