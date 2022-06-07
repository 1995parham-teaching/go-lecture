\ProvidesClass{dscription}
              [course description file]

\RequirePackage[dvipsnames]{xcolor}
\RequirePackage{booktabs}
\RequirePackage{enumerate}
\RequirePackage{minted}
\RequirePackage{enumerate}
\RequirePackage{graphicx}
\RequirePackage{keyval}

\LoadClass[]{article}

\RequirePackage[
  colorlinks=true,
  linkcolor=black,
  urlcolor=orange,
]{hyperref}
\RequirePackage[
   a4paper,
   left=2.5cm,
   right=2.5cm,
   top=2.5cm,
   bottom=2.5cm
]{geometry}
\RequirePackage[localise]{xepersian}

\usemintedstyle{one-dark}

\setlength{\parindent}{0pt}
\renewcommand{\baselinestretch}{1.5}

\eqcommand{تارنما}{href}
\eqcommand{پایان‌ساز}{makefooter}
\eqcommand{جزئیات‌درس}{coursedetails}

\date{\persiantoday}

\define@key{course}{نام}[]{\gdef\course@name{#1}}
\define@key{course}{ترم}[]{\gdef\course@semester{#1}}
\def\coursedetails#1{{\setkeys{course}{#1}}}

\renewcommand{\maketitle}{%
  \begin{center}
  \textit{بسمه تعالی}\\
  {دوره \course@name \space --- \space \course@semester}
  \\
  {پرهام الوانی}
  \\
  {\@title}
  \\
  \end{center}
}

\newcommand{\makefooter}{%
  \vfill
  \begin{center}
    این سند برپایه بسته \متن‌لاتین{\زی‌پرشین} گونه \متن‌لاتین{\گونه‌زی‌پرشین} توسعه پیدا کرده است.
    نگارش شده به تاریخ \@date
  \end{center}
}

\settextfont[
  Path = fonts/,
  UprightFont = *-Regular,
  BoldFont = *-Bold,
  ItalicFont = *-Variable
]{Vazir}

\setlatintextfont[
  Path = fonts/,
  UprightFont = *-Regular,
  BoldFont = *-Bold,
  ItalicFont = *-Italic
]{Neuton}
