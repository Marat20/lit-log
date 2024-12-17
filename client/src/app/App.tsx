import { URL } from "@/shared/api/api";
import { TG } from "@/shared/consts/consts";
import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";

export const App = () => {
  const userId = TG.initData.user.id;

  const { data } = useQuery({
    queryKey: ["user"],
    queryFn: () => fetch(`${URL}/api/user/${userId}`).then((res) => res.json()),
  });

  useEffect(() => {
    console.log(data);
  });

  return (
    <main className='main'>
      <div>Welcome to TG mini APP</div>
      {/* <Routes>
        <Route index element={<ProgressPageLazy />} />
        <Route path='/books' element={<BookListPageLazy />} />
      </Routes>
      <Footer /> */}
    </main>
  );
};
