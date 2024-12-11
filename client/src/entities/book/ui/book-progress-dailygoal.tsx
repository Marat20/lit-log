import { useQuery } from "@tanstack/react-query";
import { FC } from "react";
import { fetchBook } from "../api/api";
import cls from "./book.module.scss";

export const BookProgressDailyGoal: FC = () => {
  const { data } = useQuery({
    queryKey: ["book"],
    queryFn: fetchBook,
  });

  return (
    <section className={cls.progress_section}>
      <h2>Ежедневная цель</h2>
      <div className={cls.progress_block}>
        <label htmlFor='progress'>{data?.pagesReadToday}</label>
        <progress
          className={cls.progress}
          id='progress'
          value={data?.pagesReadToday}
          max={data?.book.dailyGoal}
        />
        <label htmlFor='progress'>{data?.book.dailyGoal}</label>
      </div>
    </section>
  );
};
