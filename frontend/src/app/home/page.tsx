"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { tweetApi } from "@/lib/api";
import { AxiosError } from "axios";
import { Tweet } from "@/api"; // 生成されたAPIクライアントから型をインポート

export default function HomePage() {
  const router = useRouter();
  const [content, setContent] = useState("");
  const [error, setError] = useState("");
  const [tweets, setTweets] = useState<Tweet[]>([]);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      router.push("/auth/login");
      return;
    }
  }, [router]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await tweetApi.tweetsPost({
        content: content,
      });
      setContent(""); // フォームをクリア
      // 投稿後にツイート一覧を再取得
      fetchTweets();
    } catch (err) {
      if (err instanceof AxiosError) {
        setError(err.response?.data?.message || "Failed to post tweet");
      }
    }
  };

  const fetchTweets = async () => {
    try {
      const response = await tweetApi.tweetsGet();
      setTweets(response.data);
    } catch (err) {
      console.error("Failed to fetch tweets:", err);
    }
  };

  useEffect(() => {
    fetchTweets();
  }, []);

  return (
    <div className="min-h-screen bg-gray-100">
      <div className="max-w-2xl mx-auto p-4">
        {/* ツイート投稿フォーム */}
        <div className="bg-white rounded-lg shadow p-4 mb-4">
          <form onSubmit={handleSubmit}>
            <textarea
              className="w-full border rounded-lg p-2 mb-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="What is happening?!"
              value={content}
              onChange={(e) => setContent(e.target.value)}
              rows={3}
            />
            {error && <div className="text-red-500 mb-2">{error}</div>}
            <div className="flex justify-end">
              <button
                type="submit"
                className="bg-blue-500 text-white px-4 py-2 rounded-full hover:bg-blue-600 transition-colors"
                disabled={!content.trim()}
              >
                Tweet
              </button>
            </div>
          </form>
        </div>

        {/* ツイート一覧 */}
        <div className="space-y-4">
          {tweets.map((tweet) => (
            <div key={tweet.id} className="bg-white rounded-lg shadow p-4">
              <div className="flex items-start space-x-3">
                <div className="flex-1">
                  <div className="flex items-center space-x-2">
                    <span className="font-bold">{tweet.user?.username}</span>
                    <span className="text-gray-500 text-sm">
                      {tweet.createdAt
                        ? new Date(tweet.createdAt).toLocaleDateString()
                        : ""}
                    </span>
                  </div>
                  <p className="mt-2">{tweet.content}</p>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
