import { Book } from "../types/book";

export const fetchBook = async (): Promise<
  { book: Book; pagesReadToday: number } | undefined
> => {
  try {
    const response = await fetch(
      "http://localhost:8080/books/YJWNnpoLOSBq6HVcK7oW9",
    );
    const result = await response.json();

    return result;
  } catch (error) {
    console.error(error);
  }
};
