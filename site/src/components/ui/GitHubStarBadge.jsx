import React, { useEffect, useState } from "react";

function GitHubIcon({ className = "w-4 h-4" }) {
  return (
    <svg
      className={className}
      viewBox="0 0 32 32"
      fill="currentColor"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path d="M15.998 0C7.164 0 0 7.192 0 16.064c-.004 3.367 1.051 6.65 3.015 9.385 1.964 2.735 4.737 4.784 7.929 5.857.8.148 1.092-.348 1.092-.774 0-.382-.014-1.392-.022-2.732-4.46.97-5.4-2.154-5.4-2.154-.726-1.856-1.776-2.35-1.776-2.35-1.454-.996.108-.976.108-.976 1.606.114 2.45 1.656 2.45 1.656 1.428 2.454 3.746 1.746 4.658 1.334.144-.958.558-1.666 1.016-2.068-3.552-.404-7.288-1.782-7.288-7.936 0-1.754.624-3.188 1.648-4.312-.166-.406-.714-2.04.154-4.25 0 0 1.344-.432 4.4 1.646a15.175 15.175 0 0 1 8.012 0c3.054-2.078 4.396-1.646 4.396-1.646.872 2.212.324 3.844.16 4.25 1.026 1.124 1.644 2.558 1.644 4.312 0 6.17-3.74 7.528-7.304 7.926.574.496 1.086 1.476 1.086 2.974 0 2.148-.02 3.88-.02 4.406 0 .43.288.93 1.1.772a16.1 16.1 0 0 0 10.86-9.86A16.063 16.063 0 0 0 32 16.064C32 7.192 24.836 0 15.998 0Z" />
    </svg>
  );
}

function StarIcon({ className = "w-4 h-4" }) {
  return (
    <svg
      className={className}
      viewBox="0 0 16 16"
      fill="currentColor"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path d="M8 .25a.75.75 0 0 1 .673.418l1.882 3.815 4.21.612a.75.75 0 0 1 .416 1.279l-3.046 2.97.719 4.192a.75.75 0 0 1-1.088.791L8 12.347l-3.766 1.98a.75.75 0 0 1-1.088-.79l.72-4.194L.818 6.374a.75.75 0 0 1 .416-1.28l4.21-.611L7.327.668A.75.75 0 0 1 8 .25Z" />
    </svg>
  );
}

export { GitHubIcon };

export default function GitHubStarBadge({ repo, className = "" }) {
  const [stars, setStars] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchStars() {
      try {
        const response = await fetch(`https://api.github.com/repos/${repo}`);
        if (response.ok) {
          const data = await response.json();
          setStars(data.stargazers_count);
        }
      } catch (error) {
        console.error("Failed to fetch GitHub stars:", error);
      } finally {
        setLoading(false);
      }
    }

    fetchStars();
  }, [repo]);

  const formatStars = (count) => {
    if (count >= 1000) {
      return `${(count / 1000).toFixed(1)}k`;
    }
    return count.toString();
  };

  return (
    <a
      href={`https://github.com/${repo}`}
      target="_blank"
      rel="noopener noreferrer"
      className={`inline-flex items-center gap-2 px-3 py-1.5 rounded-md border border-border bg-surface text-text-muted hover:text-text hover:border-primary/50 transition-all duration-200 ${className}`}
    >
      <GitHubIcon className="w-4 h-4" />
      <StarIcon className="w-3.5 h-3.5" />
      {!loading && stars !== null && (
        <>
          <span className="text-sm font-semibold">{formatStars(stars)}</span>
        </>
      )}
    </a>
  );
}
