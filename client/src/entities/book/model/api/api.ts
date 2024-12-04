import { Book } from "../types/book";

export const fetchBook = async (): Promise<{ book: Book } | undefined> => {
  try {
    const response = await fetch(
      "http://localhost:8080/books/QPzW0phCfBwuPGnW4jWvi",
    );
    const result = await response.json();
    return result;
  } catch (error) {
    console.error(error);
  }
};
