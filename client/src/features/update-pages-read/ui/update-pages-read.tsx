import { fetchBook } from "@/entities/book";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { FC, useEffect, useState } from "react";
import { fetchUpdatePagesRead, ReturnData } from "../models/api/api";
import cls from "./update-pages-read.module.scss";

export const UpdatePagesRead: FC = () => {
  const [pagesRead, setPagesRead] = useState<number>(0);
  const [id, setId] = useState<string>("");

  const queryClient = useQueryClient();

  const { data } = useQuery({
    queryKey: ["book"],
    queryFn: fetchBook,
  });

  const { mutate } = useMutation({
    mutationFn: fetchUpdatePagesRead,
    onSuccess: (data: ReturnData) => {
      queryClient.setQueryData(["book"], {
        book: data.book,
        pagesReadToday: data.pagesReadToday,
      });
    },
  });

  const onSubmit = () => {
    mutate({ pagesRead, id });
    setPagesRead(data?.book.dailyGoal ?? 0);
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
    setPagesRead(data?.book.dailyGoal ?? 0);
    setId(data?.book.ID ?? "");
  }, [data?.book.dailyGoal, data?.book.ID]);

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
