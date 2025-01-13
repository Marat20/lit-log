import { FC } from "react";
import cls from "./book.module.scss";

interface BookProgressTotalPagesProps {
  currentPage?: number;
  totalPages?: number;
}

export const BookProgressTotalPages: FC<BookProgressTotalPagesProps> = (
  props,
) => {
  const { currentPage, totalPages } = props;

  return (
    <section className={cls.progress_section}>
      <h2>Общее количество страниц</h2>
      <div className={cls.progress_block}>
        <label htmlFor='progress'>{currentPage}</label>
        <progress
          className={cls.progress}
          id='progress'
          value={currentPage}
          max={totalPages}
        />
        <label htmlFor='progress'>{totalPages}</label>
      </div>
    </section>
  );
};
