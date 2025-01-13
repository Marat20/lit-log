import {
  BookProgressDailyGoal,
  BookProgressTotalPages,
  BookTitle,
  fetchBook,
  init,
} from "@/entities/book";
import { UpdatePagesRead } from "@/features/update-pages-read";
import { useQuery } from "@tanstack/react-query";
import { FC } from "react";

const ProgressPage: FC = () => {
  // const navigate = useNavigate();

  const { data: bookId } = useQuery({
    queryKey: ["bookId"],
    queryFn: init,
  });

  const { data, isLoading } = useQuery({
    queryKey: ["book", bookId],
    queryFn: fetchBook,
    enabled: !!bookId,
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (data?.error) {
    return <div>Error</div>;
  }

  // if (data?.error) {
  //   navigate("/new");
  // }

  return (
    <>
      <BookTitle
        title={data?.currentBook?.title}
        author={data?.currentBook?.author}
      />
      <BookProgressTotalPages
        currentPage={data?.currentBook?.currentPage}
        totalPages={data?.currentBook?.totalPages}
      />
      <BookProgressDailyGoal
        pagesReadToday={data?.pagesReadToday}
        dailyGoal={data?.currentBook?.dailyGoal}
      />
      <UpdatePagesRead
        dailyGoal={data?.currentBook?.dailyGoal}
        bookId={data?.currentBook?.id}
      />
    </>
  );
};

export default ProgressPage;
