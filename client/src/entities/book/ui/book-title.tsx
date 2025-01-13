import { FC } from "react";

interface BookTitleProps {
  title?: string;
  author?: string;
}

export const BookTitle: FC<BookTitleProps> = (props) => {
  const { title, author } = props;

  return (
    <section>
      <h1>{title}</h1>
      <h2>{author}</h2>
    </section>
  );
};
