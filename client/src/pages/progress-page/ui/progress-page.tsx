import {
  BookProgressDailyGoal,
  BookProgressTotalPages,
  BookTitle,
  fetchBook,
} from "@/entities/book";
import { UpdatePagesRead } from "@/features/update-pages-read";
import { useQuery } from "@tanstack/react-query";
import { FC } from "react";
import { useParams } from "react-router";

const ProgressPage: FC = () => {
  const params = useParams<{ bookId: string }>();

  const { data, isLoading } = useQuery({
    queryKey: ["book"],
    gcTime: 1,
    queryFn: () => fetchBook(params.bookId),
    enabled: !!params.bookId,
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (data?.error) {
    return <div>Error</div>;
  }

  return (
    <>
      <BookTitle title={data?.book?.title} author={data?.book?.author} />
      <BookProgressTotalPages
        currentPage={data?.book?.currentPage}
        totalPages={data?.book?.totalPages}
      />
      <BookProgressDailyGoal
        pagesReadToday={data?.pagesReadToday}
        dailyGoal={data?.book?.dailyGoal}
      />
      <UpdatePagesRead
        dailyGoal={data?.book?.dailyGoal}
        bookId={data?.book?.id}
      />
    </>
  );
};

export default ProgressPage;
