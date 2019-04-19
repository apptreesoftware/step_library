
using System;

namespace Famis.Model
{
    public class UpsertResponse<T>
    {
        public bool Success { get; set; }
        public String Message { get; set; }
        public T Object { get; set; }

        public UpsertResponse(bool success, string message, T o) {
            Success = success;
            Message = message;
            Object = o;
        }
    }
}