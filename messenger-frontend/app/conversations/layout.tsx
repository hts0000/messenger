import React from "react";

export default function ConversationsLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div>
      Conversations Layout
      {children}
    </div>
  );
}
