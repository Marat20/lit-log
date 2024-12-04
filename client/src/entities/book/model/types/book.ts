export interface Book {
  ID: string;
  isActive: boolean;
  isDone: boolean;
  createdAt: string;
  updatedAt: string;
  finishedAt: string;
  title: string;
  author: string;
  totalPages: number;
  currentPage: number;
  dailyGoal: number;
  pagesRead: number;
}
