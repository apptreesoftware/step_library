using System;
using Newtonsoft.Json;

namespace Famis.Model {
    public class AuthInfo {
        [JsonProperty("access_token")] public string AccessToken { get; set; }
        [JsonProperty("token_type")] public string TokenType { get; set; }
        [JsonProperty("expires_in")] public int ExpiresIn { get; set; }
        [JsonProperty("refresh_token")] public string RefreshToken { get; set; }
        [JsonProperty("user_id")] public string UserId { get; set; }
        [JsonProperty("first_name")] public string FirstName { get; set; }
        [JsonProperty("last_name")] public string LastName { get; set; }
        [JsonProperty(".expires")] public DateTime Expires { get; set; }
    }

    public class AuthResponse {
        public AuthInfo Item { get; set; }
        public bool Result { get; set; }
        public int Context { get; set; }
        public string Message { get; set; }

        public bool IsExpired => Item.Expires.Ticks < DateTime.Now.Ticks;
    }
}