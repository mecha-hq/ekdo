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
          <img src="trivy-logo.png" alt="Trivy Logo" />
        </div>
      </div>
      <h1 class="text-2xl font-bold ml-4 place-self-center">
        <p>
          Trivy Scan Report
        </p>
      </h1>
    </div>

    <!-- Docker Image Information -->
    <div class="mt-8">
      <div class="overflow-x-clip">
        <h2 class="text-xl font-bold">Images information</h2>

        <table class="min-w-full border rounded mt-4">
          <tbody class="divide-none">
            <tr style="border: none">
              <td class="font-bold py-2 px-4 border-b">Artifact Name</td>
              <td class="py-2 px-4 border-b">{{ title .ArtifactName }}</td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">Artifact Type</td>
              <td class="py-2 px-4 border-b">{{ title (.ArtifactType | toString) }}</td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">OS Kind</td>
              <td class="py-2 px-4 border-b">{{ title .Metadata.ImageConfig.OS }}</td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">OS Family</td>
              <td class="py-2 px-4 border-b">{{ title (.Metadata.OS.Family | toString) }}</td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">OS Name</td>
              <td class="py-2 px-4 border-b">{{ title .Metadata.OS.Name }}</td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">Architecture</td>
              <td class="py-2 px-4 border-b">{{ title .Metadata.ImageConfig.Architecture }}</td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">Env Vars</td>
              <td class="py-2 px-4 border-b">
                {{ range $env := .Metadata.ImageConfig.Config.Env }}
                  {{ $env }}<br />
                {{ end }}
              </td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">Entrypoint</td>
              <td class="py-2 px-4 border-b">
                {{ range $entrypoint := .Metadata.ImageConfig.Config.Entrypoint }}
                  {{ $entrypoint }}&nbsp;
                {{ end }}
              </td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">Cmd</td>
              <td class="py-2 px-4 border-b">
                {{ range $cmd := .Metadata.ImageConfig.Config.Cmd }}
                  {{ $cmd }}&nbsp;
                {{ end }}
              </td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">User</td>
              <td class="py-2 px-4 border-b">{{ .Metadata.ImageConfig.Config.User }}</td>
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
            <th class="py-2 px-4 border-b">Target</th>
            <th class="py-2 px-4 border-b">Class</th>
            <th class="py-2 px-4 border-b">Type</th>
          </tr>
        </thead>
        <tbody>
          {{ range $result := .Results }}
          <tr>
            <td class="py-2 px-4 border-b">{{ $result.Target }}</td>
            <td class="py-2 px-4 border-b">{{ $result.Class }}</td>
            <td class="py-2 px-4 border-b">{{ $result.Type }}</td>
          </tr>
          {{ end }}
          <!-- Add more rows if there are multiple vulnerabilities -->
        </tbody>
      </table>
    </div>
  </div>
</body>

</html>
