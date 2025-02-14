import { fetchAllAgencies } from "ecfr-analyzer/service/EcfrService";
import Error from "ecfr-analyzer/components/Error";

export default async function Page() {
  const titlesResponse = await fetchAllAgencies();

  if (titlesResponse.err) {
    return <Error message={titlesResponse.err.message} />;
  }

  // console.log(titlesResponse.data.agencies[45]);

  return (
    <table>
      <tbody>
        {titlesResponse?.data?.agencies?.map((it, i) => (
          <tr key={i}>
            <td>{it.slug}</td>
            <td>{it.name}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
