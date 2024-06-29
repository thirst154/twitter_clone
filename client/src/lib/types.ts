export interface GetPostsData {
  data: Post[];
}

export interface Post {
  id: number;
  text: string;
  date_published: Date;
  user: User;
}

export interface User{
  first_name: string;
  last_name: string;
  username: string;
}
