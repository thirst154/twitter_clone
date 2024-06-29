import { Card, CardContent, CardHeader } from "./ui/card";

export default function Message(){

  return (
    <Card className="flex flex-col">
      <CardHeader>
      <div className="flex max-w-fit gap-2 items-center">
        <span className="text-lg text-primary font-semibold">Thomas Hirst</span>
        <span className="text-sm text-muted-foreground">@username</span>
      </div>
      </CardHeader>
      <CardContent>
      <span className="text-sm">Lorem ipsum dolor sit, amet consectetur adipisicing elit. A minus consequatur culpa officiis provident dolor recusandae, explicabo quibusdam optio animi corrupti alias cupiditate iure ipsa ratione voluptate itaque non quisquam!</span>
      </CardContent>
    </Card>
  )
}