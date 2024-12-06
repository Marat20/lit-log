import { useQuery } from "@tanstack/react-query";
import { FC } from "react";
import { fetchBook } from "../model/api/api";
import cls from "./book.module.scss";

export const BookProgressTotalPages: FC = () => {
  const { data } = useQuery({
    queryKey: ["book"],
    queryFn: fetchBook,
  });

  return (
    <section className={cls.progress_section}>
      <h2>Общее количество страниц</h2>
      <div className={cls.progress_block}>
        <label htmlFor='progress'>{data?.book.currentPage}</label>
        <progress
          className={cls.progress}
          id='progress'
          value={data?.book.currentPage}
          max={data?.book.totalPages}
        />
        <label htmlFor='progress'>{data?.book.totalPages}</label>
      </div>
    </section>
  );
};
