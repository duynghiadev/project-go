import SinglePost from "@/components/SinglePost";

const page = ({ params }: any) => {
  const { id } = params;

  console.log(id);

  return (
    <div>
      <SinglePost id={id} />
    </div>
  );
};

export default page;
