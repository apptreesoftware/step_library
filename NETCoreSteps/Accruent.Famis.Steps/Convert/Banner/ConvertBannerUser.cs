using System;
using Banner;
using Famis.Model;
using StepCore;

namespace Accruent.Famis.Steps.Convert.Banner {
    [StepDescription("convert_banner_user", Description = "Converts a Banner Requester to a Famis User.")]
    public class ConvertBannerUser : Step {
        [Input(Description = "The banner requester to convert")] 
        public Requestor BannerUser { get; set; }
        [Input(Description = "The default profile id to apply to this user. This is a required field in FAMIS.")] 
        public int ProfileId { get; set; }
        [Input(Description = "The default State id to apply to this user. This is a required field in FAMIS")] 
        public int StateId { get; set; }
        [Input(Description = "The default Country id to apply to this user. This is a required field in FAMIS")] 
        public int CountryId { get; set; }
        [Input(Description = "The Active status of this user in FAMIS. This is a required field in FAMIS")] 
        public bool ActiveFlag { get; set; }
        [Input(Description = "The default Company id to apply to this user. This is a required field in FAMIS")] 
        public int CompanyId { get; set; }
        [Input(Description = "The default Street address to apply to this user. This is a required field in FAMIS")] 
        public string Addr1 { get; set; }
        [Input(Description = "The default City to apply to this user. This is a required field in FAMIS")] 
        public string City { get; set; }
        [Input(Description = "The default Zip to apply to this user. This is a required field in FAMIS")]
        public string Zip { get; set; }

        
        [Output(Description = "The converted famis user")] public User FamisUser { get; set; }

        public override void Execute() {
            FamisUser = new User {
                FirstName = BannerUser.FirstName,
                LastName = BannerUser.LastName,
                UserName = BannerUser.Id,
                BusPhone = BannerUser.Id,
                Title = BannerUser.Department,
                ActiveFlag = true,
                StateId = StateId,
                CountryId = CountryId,
                ProfileId = ProfileId,
                CompanyId = CompanyId,
                Addr1 = Addr1,
                City = City,
                Zip = Zip,
                Password = BannerUser.FirstName + BannerUser.LastName,
                ExternalId = BannerUser.Id
            };
        }
    }
}