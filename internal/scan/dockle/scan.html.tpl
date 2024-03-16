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
          <img src="dockle-logo.png" alt="Dockle Logo" />
        </div>
      </div>
      <h1 class="text-2xl font-bold ml-4 place-self-center">
        <p>
          Dockle Scan Report
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
              <td class="py-2 px-4 border-b">{{ title .ImageName }}</td>
            </tr>
            <tr>
              <td class="font-bold py-2 px-4 border-b">Artifact Type</td>
              <td class="py-2 px-4 border-b">container image</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Details Table -->
    <div class="mt-8 mb-4">
      <h2 class="text-xl font-bold">Details</h2>

      <table class="min-w-full border rounded mt-4">
        <colgroup>
          <col span="1" style="width: 20%;">
          <col span="1" style="width: 20%;">
          <col span="1" style="width: 10%;">
          <col span="1" style="width: 50%;">
        </colgroup>
        <thead class="bg-gray-200">
          <tr>
            <th class="py-2 px-4 border-b">Code</th>
            <th class="py-2 px-4 border-b">Title</th>
            <th class="py-2 px-4 border-b">Level</th>
            <th class="py-2 px-4 border-b">Alerts</th>
          </tr>
        </thead>
        <tbody>
          {{ range .Details }}
          <tr>
            <td class="py-2 px-4 border-b">{{ .Code }}</td>
            <td class="py-2 px-4 border-b">{{ .Title }}</td>
            <td class="py-2 px-4 border-b">{{ .Level }}</td>
            <td class="py-2 px-4 border-b">
            {{- range .Alerts }}
              <p>{{ . }}</p>
            {{- end }}
            </td>
          </tr>
          {{ end }}
        </tbody>
      </table>
    </div>
  </div>
</body>

</html>