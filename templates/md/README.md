{{ define "readme" }}

<!-- PROJECT LOGO -->
<div align="center">
<a href="https://github.com/jordanopensource/nuha-web">
  <img src=".github/images/logo.svg" alt="Logo" width="80" height="80">
</a>

### {{ .Title }}

{{ .BriefDescription }}

[Explore the docs »]({{ .DocsUrl }})

.
[Report Bug]({{ .IssuesUrl }})
.
[Request Feature]({{ .IssuesUrl }})

[![Build Status]({{ .BuildBadge }})]({{ .BuildStatusUrl }})

</div>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#running">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

{{ .Description }}

### Built With 🤖

{{ range .BuiltWith }}

- [{{ .Title }}]({{ .Link }})
  {{ end }}

---

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

{{ .PrerequisitesMd }}

### Installation

{{ .InstallationMd }}

### Running

{{ .RunningMd }}

---

<!-- ROADMAP -->

## Roadmap

See the [open issues]({{ .IssuesUrl }}) for a list of proposed features (and known issues).

---

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/amazing-feature`)
3. Commit your Changes (`git commit -m 'Add some amazing-feature'`)
4. Push to the Branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

<!-- LICENSE -->

## License

Distributed under the Apache License 2.0. See [LICENSE](LICENSE) for more information.

---

<!-- CONTACT -->

## Contact

Jordan Open Source Association - [@jo_osa](https://twitter.com/@jo_osa)

Project Link: [https://github.com/jordanopensource/nuha-web](https://github.com/jordanopensource/nuha-web)
{{ end }}
