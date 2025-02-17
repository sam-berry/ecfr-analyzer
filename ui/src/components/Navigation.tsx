export default function Navigation({
  title = "Code of Federal Regulations",
}: {
  title?: string;
}) {
  return (
    <>
      <div className="font-title bg-primary text-light fixed left-0 right-0 top-0 z-[100] px-4 py-2 text-center text-2xl uppercase tracking-wide md:text-3xl">
        {title}
      </div>
      <div className="h-[3rem] w-full md:h-[3.3rem]" />
    </>
  );
}
