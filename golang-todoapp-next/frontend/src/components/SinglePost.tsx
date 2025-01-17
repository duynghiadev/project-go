import DeleteTodo from "@/components/DeleteTodo";
import EditTodo from "@/components/EditTodo";

const getSingleData = async (id: string) => {
  const res = await fetch(`http://localhost:8000/todos/${id}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    throw new Error("Failed to get a single post todo");
  }
  return res.json();
};

const SinglePost = async ({ id }: { id: string }) => {
  const todo = await getSingleData(id);

  return (
    <div className="flex items-center justify-center">
      {todo && (
        <div className="h-96 bg-yellow-400 p-12 rounded">
          <p className="text-3xl uppercase">Title: {todo.Title}</p>
          <p>Content: {todo.Content}</p>
          <p>Createdat:{todo.Createdat}</p>

          <div className="m-5">
            <DeleteTodo todo={todo} />
          </div>

          <div className="m-5">
            <EditTodo todo={todo} />
          </div>
        </div>
      )}
    </div>
  );
};

export default SinglePost;
