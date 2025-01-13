import { useMutation, useQueryClient } from "@tanstack/react-query";
import { FC, useEffect, useState } from "react";
import { fetchUpdatePagesRead, ReturnData } from "../models/api/api";
import cls from "./update-pages-read.module.scss";

interface UpdatePagesReadProps {
  dailyGoal?: number;
  bookId?: string;
}

export const UpdatePagesRead: FC<UpdatePagesReadProps> = (props) => {
  const { dailyGoal, bookId } = props;
  const [pagesRead, setPagesRead] = useState<number>(0);

  const queryClient = useQueryClient();

  const { mutate } = useMutation({
    mutationFn: fetchUpdatePagesRead,
    onSuccess: (data: ReturnData) => {
      queryClient.setQueryData(["book"], {
        currentBook: data.currentBook,
        pagesReadToday: data.pagesReadToday,
      });
      queryClient.setQueryData(["bookId"], {
        bookId: data.currentBook.id
      });
    },
  });

  const onSubmit = () => {
    if (!bookId) {
      return;
    }
    mutate({ pagesRead, bookId });
    setPagesRead(dailyGoal ?? 0);
  };

  const handleIncrement = () => {
    setPagesRead((prev) => prev + 1);
  };

  const handleDecrement = () => {
    if (pagesRead > 1) {
      setPagesRead((prev) => prev - 1);
    }
  };

  useEffect(() => {
    setPagesRead(dailyGoal ?? 0);
  }, [dailyGoal]);

  return (
    <section>
      <div className={cls.fetch_btn_block}>
        <button className={cls.fetch_btn} onClick={onSubmit}>
          {pagesRead}
        </button>
      </div>
      <div className={cls.counter_btn_block}>
        <button className={cls.counter_btn} onClick={handleDecrement}>
          -
        </button>
        <button className={cls.counter_btn} onClick={handleIncrement}>
          +
        </button>
      </div>
    </section>
  );
};
