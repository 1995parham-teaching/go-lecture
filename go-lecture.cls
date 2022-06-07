\ProvidesClass{go-lecture}
              [Parham Alvani's theme for Go programming course]

\RequirePackage[dvipsnames]{xcolor}
\RequirePackage{booktabs}
\RequirePackage{minted}

\LoadClass[]{beamer}

\usemintedstyle{one-dark}

\usetheme[secheader]{Boadilla}
\usecolortheme{crane}

\author[Parham Alvani]{%
  Parham Alvani\\
  \texttt{parham.alvani@gmail.com}
}
\title{Go Programming and Backend Development}
\institute[Mapsa]{Mapsa HR}
\date{\Today}

\hypersetup{colorlinks=true,allcolors=.,urlcolor=orange}
\setbeamertemplate{headline}[infolines theme]
% the following command suppresses all navigation symbols
\setbeamertemplate{navigation symbols}{}

\AtBeginSection[]
{%
  \begin{frame}
    \frametitle{Outline}
    \tableofcontents[currentsection]
  \end{frame}
}
