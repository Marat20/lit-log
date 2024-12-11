import clsx from 'clsx';
import { FC } from 'react';
import { Book } from '../model/types/book';
import cls from './book.module.scss';

interface BookCardProps {
  bookData: Book;
}

export const BookCard: FC<BookCardProps> = (props) => {
  const { title, author, isDone, isActive, totalPages, currentPage } =
    props.bookData;

  const classes = {
    [cls.isDone]: isDone,
    [cls.isActive]: isActive,
  };

  return (
    <div className={clsx('card', classes)}>
      <h3>{title}</h3>
      <div className={cls.down}>
        <h4>{author}</h4>
        <div>
          {currentPage}/{totalPages}
        </div>
      </div>
    </div>
  );
};
