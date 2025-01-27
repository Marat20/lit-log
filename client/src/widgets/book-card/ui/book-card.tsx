import { Book } from "@/entities/book";
import { DeleteBook } from "@/features/delete-book";
import clsx from "clsx";
import { FC } from "react";
import { Link } from "react-router";
import cls from "./book-card.module.scss";

interface BookCardProps {
  bookData: Book;
}

export const BookCard: FC<BookCardProps> = (props) => {
  const { title, author, isDone, isActive, totalPages, currentPage, id } =
    props.bookData;

  const classes = {
    [cls.isDone]: isDone,
    [cls.isActive]: isActive,
  };

  return (
    <div className={clsx(cls.card, { ...classes })}>
      <Link key={id} to={`/${id}`} className={cls.card__link}>
        <div className={cls.card__info}>
          <h3>{title}</h3>
          <h4>{author}</h4>
        </div>
        <div className={cls.card__pages}>
          {currentPage}/{totalPages}
        </div>
      </Link>
      <DeleteBook bookId={id} />
    </div>
  );
};
