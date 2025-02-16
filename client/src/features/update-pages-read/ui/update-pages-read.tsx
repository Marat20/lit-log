import { FetchBookReturnData } from "@/entities/book/api/api";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { FC, useEffect, useState } from "react";
import { fetchUpdatePagesRead } from "../api/api";
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
    onMutate: async (newData) => {
      await queryClient.cancelQueries({ queryKey: ["book"] });
      const previousData = queryClient.getQueryData<FetchBookReturnData>([
        "book",
      ]);

      queryClient.setQueryData(["book"], (old: FetchBookReturnData) => ({
        ...old,
        book: {
          ...old.book,
          currentPage: old.book?.currentPage || 0 + pagesRead,
        },
        pagesReadToday: (old?.pagesReadToday || 0) + newData.pagesRead,
      }));

      return { previousData };
    },

    onSuccess: (data) => {
      queryClient.setQueryData(["book"], {
        book: data.currentBook,
        pagesReadToday: data.pagesReadToday,
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
        <button className={cls.counter_btn} onClick={handleDecrement}>
          -
        </button>
        <button className={cls.fetch_btn} onClick={onSubmit}>
          {pagesRead}
        </button>
        <button className={cls.counter_btn} onClick={handleIncrement}>
          +
        </button>
      </div>
    </section>
  );
};
