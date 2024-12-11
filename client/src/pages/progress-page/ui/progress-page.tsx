import {
  BookProgressDailyGoal,
  BookProgressTotalPages,
  BookTitle,
} from '@/entities/book';
import { UpdatePagesRead } from '@/features/update-pages-read';
import { FC } from 'react';

const ProgressPage: FC = () => {
  return (
    <>
      <BookTitle />
      <BookProgressTotalPages />
      <BookProgressDailyGoal />
      <UpdatePagesRead />
    </>
  );
};

export default ProgressPage;
