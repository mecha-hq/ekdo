<html lang="en" data-theme="dark">

<head>
  <link href="https://cdn.jsdelivr.net/npm/daisyui@4.7.2/dist/full.min.css" rel="stylesheet" type="text/css" />
  <script src="https://cdn.tailwindcss.com"></script>
</head>

<body>
  <div class="max-w-4xl mx-auto bg-white p-8 rounded shadow">
    <div class="flex flex-row">
      <div class="avatar place-self-center">
        <div class="mask mask-squircle w-12 h-12">
          <img src="grype-logo.png" alt="Grype Logo" />
        </div>
      </div>
      <h1 class="text-2xl font-bold ml-4 place-self-center">
        <p>
          Grype Scan Report
        </p>
      </h1>
    </div>

    <!-- Docker Image Information -->
    <div class="mt-8">
      <div class="overflow-x-clip">
        <h2 class="text-xl font-bold">Information</h2>

        <table class="min-w-full border rounded mt-4">
          <tbody class="divide-none">
            <tr style="border: none">
              <td class="font-bold py-2 px-4 border-b">Artifact Name</td>
              <td class="py-2 px-4 border-b">{{ title (index .Source.Target "userInput") }}</td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">Artifact Type</td>
              <td class="py-2 px-4 border-b">{{ title .Source.Type }}</td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">OS Kind</td>
              <td class="py-2 px-4 border-b">{{ title (index .Source.Target "os") }}</td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">OS Name</td>
              <td class="py-2 px-4 border-b">{{ title (print .Distro.Name " " .Distro.Version) }}</td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">Architecture</td>
              <td class="py-2 px-4 border-b">{{ title (index .Source.Target "architecture") }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Vulnerabilities Table -->
    <div class="mt-8 mb-4">
      <h2 class="text-xl font-bold">Vulnerabilities</h2>

      <table class="min-w-full border rounded mt-4">
        <thead class="bg-gray-200">
          <tr>
            <th class="py-2 px-4 border-b">Id</th>
            <th class="py-2 px-4 border-b">Severity</th>
            <th class="py-2 px-4 border-b">Fix state</th>
            <th class="py-2 px-4 border-b">Description</th>
          </tr>
        </thead>
        <tbody>
          {{ range .Matches }}
          <tr>
            <td class="py-2 px-4 border-b">
              <a class="link" href="{{ .Vulnerability.DataSource }}">{{ .Vulnerability.ID }}</a>
            </td>
            <td class="py-2 px-4 border-b">{{ .Vulnerability.Severity }}</td>
            <td class="py-2 px-4 border-b">{{ .Vulnerability.Fix.State }}</td>
            <td class="py-2 px-4 border-b">{{ .Vulnerability.Description }}</td>
          </tr>
          {{ end }}
        </tbody>
      </table>
    </div>
  </div>
</body>

</html>
