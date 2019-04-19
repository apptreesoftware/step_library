using StepCore;

namespace Accruent.Famis.Steps {
    public abstract class ServiceStep : StepAsync {
        [Input(Description = "FAMIS service url")]
        public string Url { get; set; }
        [Input(Description = "FAMIS service username")]
        public string Username { get; set; }
        [Input(Description = "FAMIS service password")]
        public string Password { get; set; }
    }
}