<script lang="ts">
  export let data;

</script>

<section id="tables">
  <h2>Hi, {data.name}!</h2>
  <figure>
    <table role="grid">
      <thead>
        <tr>
          <th scope="col">#</th>
          <th scope="col">Title</th>
          <th scope="col">Release Year</th>
        </tr>
      </thead>
      <tbody>
        {#each data.films as film, index}
        <tr>
          <th scope="row">{index}</th>
          <td>{film.title}</td>
          <td>{film.releaseYear}</td>
        </tr>
        {/each}
      </tbody>
    </table>
  </figure>
</section>
