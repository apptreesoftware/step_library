using System;
using System.Threading.Tasks;
using StepCore;

namespace Accruent.Famis.Steps {
   
    [PackageDefinition("Accruent.Famis")]
    class Program {
        static async Task Main(string[] args) {
            await PackageManager.Run(args);
        }
    }
}