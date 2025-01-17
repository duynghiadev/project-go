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
        <div>
          <p className="text-3xl uppercase">Title: {todo.Title}</p>
          <p>Content: {todo.Content}</p>
          <p>Createdat:{todo.Createdat}</p>
        </div>
      )}
    </div>
  );
};

export default SinglePost;
