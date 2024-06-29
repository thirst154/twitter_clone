import Link from "next/link";

export default function NavBar(){
  return (
    <header className="sticky top-0 flex h-16 items-center content-between gap-4 border-b bg-background px-4 md:px-6">
     <nav className="flex h-full items-center gap-4">
      <h3>Website</h3>
      <Link href="/" className="text-muted-foreground transition-colors hover:text-foreground">Home</Link>
      <Link href="/messages/" className="text-muted-foreground transition-colors hover:text-foreground">Messages</Link>
      <Link href="/about" className="text-muted-foreground transition-colors hover:text-foreground">About</Link>
     </nav>
     <nav className="flex items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
      <Link href="/login" className="text-primary font-semibold">Login</Link>
      <Link href="/signup" className="text-primary font-semibold">Sign Up</Link>
     </nav>
    </header>
  )
}
