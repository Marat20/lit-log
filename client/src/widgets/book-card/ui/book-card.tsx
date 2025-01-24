import { Book } from "@/entities/book";
import { DeleteBook } from "@/features/delete-book";
import clsx from "clsx";
import { FC } from "react";
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
      <div>
        <h3>{title}</h3>
        <div className={cls.down}>
          <h4>{author}</h4>
          <div>
            {currentPage}/{totalPages}
          </div>
        </div>
      </div>
      <DeleteBook bookId={id} className={cls.btn_delete} />
    </div>
  );
};
