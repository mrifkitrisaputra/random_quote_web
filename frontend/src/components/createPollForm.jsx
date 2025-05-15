import React, { useState } from "react";
import { createPoll } from "../service/api";

const CreatePollForm = ({ onCreate }) => {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [options, setOptions] = useState([""]);
  const [expiresAt, setExpiresAt] = useState("2025-12-31T23:59");
  const [allowRealtime, setAllowRealtime] = useState(false);

  const handleAddOption = () => {
    setOptions([...options, ""]);
  };

  const handleOptionChange = (e, index) => {
    const newOptions = [...options];
    newOptions[index] = e.target.value;
    setOptions(newOptions);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Validasi jumlah opsi
    if (options.filter((opt) => opt.trim() !== "").length < 2) {
      alert("Minimal 2 opsi jawaban");
      return;
    }

    // Validasi title
    if (!title.trim()) {
      alert("Judul polling harus diisi");
      return;
    }

    // Konversi expiresAt dari datetime-local ("2025-12-31T23:59") → jadi "2025-12-31 23:59:00"
    const formattedExpiresAt = expiresAt
      ? expiresAt.replace("T", " ") + ":00"
      : "";

    const payload = {
      title,
      description,
      options: options.filter((opt) => opt.trim() !== ""),
      expires_at: formattedExpiresAt,
      allow_realtime_results: allowRealtime,
    };

    try {
      const data = await createPoll(payload);
      onCreate(data.poll_id, data.url);
    } catch (err) {
      alert(err.message || "Gagal membuat polling");
      console.error(err);
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="max-w-lg mx-auto p-4 bg-white shadow rounded"
    >
      <h2 className="text-xl font-bold mb-4">Buat Polling Baru</h2>

      <input
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="Judul Polling"
        className="w-full p-2 mb-2 border rounded"
        required
      />

      <textarea
        value={description}
        onChange={(e) => setDescription(e.target.value)}
        placeholder="Deskripsi polling"
        className="w-full p-2 mb-2 border rounded"
      />

      {options.map((opt, idx) => (
        <input
          key={idx}
          value={opt}
          onChange={(e) => handleOptionChange(e, idx)}
          placeholder={`Opsi ${idx + 1}`}
          className="w-full p-2 mb-2 border rounded"
          required
        />
      ))}

      <button
        type="button"
        onClick={handleAddOption}
        className="mb-4 text-blue-600 underline"
      >
        Tambah Opsi
      </button>

      <label className="flex items-center gap-2 mb-4">
        <input
          type="checkbox"
          checked={allowRealtime}
          onChange={() => setAllowRealtime(!allowRealtime)}
        />
        Izinkan lihat hasil saat aktif
      </label>

      <input
        type="datetime-local"
        value={expiresAt}
        onChange={(e) => setExpiresAt(e.target.value)}
        className="w-full p-2 my-2 border rounded"
      />

      <button
        type="submit"
        className="bg-blue-600 text-white px-4 py-2 w-full rounded hover:bg-blue-700"
      >
        Buat Polling
      </button>
    </form>
  );
};

export default CreatePollForm;
