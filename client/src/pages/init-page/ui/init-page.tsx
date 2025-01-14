import { init } from "@/entities/book";
import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";
import { useNavigate } from "react-router";

export const InitPage = () => {
  const { data } = useQuery({
    queryKey: ["bookId"],
    queryFn: init,
  });

  const navigate = useNavigate();

  useEffect(() => {
    if (data?.bookId) {
      navigate(`/${data.bookId}`);
    } else {
      navigate("/new");
    }
  });

  return <></>;
};
