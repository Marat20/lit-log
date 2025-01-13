import { FC } from "react";
import cls from "./book.module.scss";

interface BookProgressDailyGoalProps {
  pagesReadToday?: number;
  dailyGoal?: number;
}

export const BookProgressDailyGoal: FC<BookProgressDailyGoalProps> = (
  props,
) => {
  const { pagesReadToday, dailyGoal } = props;

  return (
    <section className={cls.progress_section}>
      <h2>Ежедневная цель</h2>
      <div className={cls.progress_block}>
        <label htmlFor='progress'>{pagesReadToday}</label>
        <progress
          className={cls.progress}
          id='progress'
          value={pagesReadToday}
          max={dailyGoal}
        />
        <label htmlFor='progress'>{dailyGoal}</label>
      </div>
    </section>
  );
};
