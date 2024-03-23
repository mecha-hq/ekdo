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
        <h2 class="text-xl font-bold">Information</h2>

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
              <td class="font-bold py-2 px-4 border-b">OS Name</td>
              <td class="py-2 px-4 border-b">{{ title (print (.Metadata.OS.Family | toString) " " .Metadata.OS.Name) }}</td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">Architecture</td>
              <td class="py-2 px-4 border-b">{{ title .Metadata.ImageConfig.Architecture }}</td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">Env Vars</td>
              <td class="py-2 px-4 border-b">
                <table>
                {{ $lastEnv := last .Metadata.ImageConfig.Config.Env }}
                {{ range $i, $env := .Metadata.ImageConfig.Config.Env }}

                  {{ $borderClass := "border-b" }}
                  {{ if eq $env $lastEnv }}
                    {{ $borderClass = "border-b-0" }}
                  {{ end }}

                  {{ $envParts := splitList "=" $env }}
                  <tr {{ if eq $env $lastEnv }}class="border-b-0"{{ end }}>
                    <td class="py-2 px-4 font-bold {{ $borderClass }}">
                      {{ index $envParts 0 }}
                    </td>
                    <td class="py-2 px-4 mx-auto break-all {{ $borderClass }}">
                      {{ index $envParts 1 }}
                    </td>
                  </tr>
                {{ end }}
                </table>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Vulnerabilities Table -->
    <div class="mt-8 mb-4">
      <h2 class="text-xl font-bold">Vulnerabilities</h2>
        {{ range $result := .Results }}
        <table class="min-w-full border rounded mt-4">
          <colgroup>
            <col span="1" style="width: 20%;">
            <col span="1" style="width: 10%;">
            <col span="1" style="width: 10%;">
            <col span="1" style="width: 60%;">
          </colgroup>
          <thead class="bg-gray-200">
            <tr>
              <th class="py-2 px-4 border-b uppercase" colspan="4">
                {{ $result.Class }}({{ $result.Type }})
              </th>
            </tr>
            <tr>
              <th class="py-2 px-4 border-b">Id</th>
              <th class="py-2 px-4 border-b">Severity</th>
              <th class="py-2 px-4 border-b">Status</th>
              <th class="py-2 px-4 border-b">Description</th>
            </tr>
          </thead>
          <tbody>
          {{ range $vuln := $result.Vulnerabilities }}
          <tr>
            <td class="py-2 px-4 border-b">
              <a class="link" href="{{ $vuln.PrimaryURL }}">{{ $vuln.VulnerabilityID }}</a>
            </td>
            <td class="py-2 px-4 border-b">{{ title (lower $vuln.Severity) }}</td>
            <td class="py-2 px-4 border-b">{{ title (lower $vuln.Status.String) }}</td>
            <td class="py-2 px-4 border-b">{{ $vuln.Description }}</td>
          </tr>
          {{ end }}
          </tbody>
        </table>
        {{ end }}
    </div>
  </div>
</body>

</html>
