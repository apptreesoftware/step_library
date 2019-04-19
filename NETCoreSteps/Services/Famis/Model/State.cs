namespace Famis.Model
{
    public class State
    {
        public string Id { get; set; }
        public string CountryId { get; set; }
        public string CountryName { get; set; }
        public string Name { get; set; }
        public bool ActiveFlag { get; set; }
        public string Description { get; set; }
        public string Abbreviation { get; set; }
        public string UpdateDate { get; set; }
        public string UpdatedByName { get; set; }
        public string TabOrder { get; set; }
        public bool DefaultFlag { get; set; }

    }
}