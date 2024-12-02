export interface Book {
  id: string;
  createdAt: string;
  updatedAt: string;
  title: string;
  author: string;
  totalPages: number;
  currentPage: number;
  dailyGoal: number;
}
